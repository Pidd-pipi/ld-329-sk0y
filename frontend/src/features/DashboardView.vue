<template>
  <main class="page-shell" v-loading="loading">
    <AppHeader :unread="overview?.metrics.unread ?? 0" />

    <section v-if="overview" class="metrics-grid">
      <MetricCard label="已发布技能" :value="overview.metrics.skills" />
      <MetricCard label="活跃需求" :value="overview.metrics.needs" />
      <MetricCard label="智能匹配" :value="overview.metrics.matches" />
      <MetricCard label="评价记录" :value="overview.metrics.reviews" />
    </section>

    <el-alert v-if="error" :title="error" type="error" show-icon />

    <section v-if="overview" class="workspace-grid">
      <div class="panel">
        <h2>技能发布</h2>
        <FeatureCard v-for="skill in overview.skills" :key="skill.id" :title="skill.title" :description="skill.description">
          <template #tag><el-tag>{{ skill.category }} {{ skill.level }}%</el-tag></template>
          <div class="tag-row">
            <el-tag v-for="slot in skill.timeSlots" :key="slot" effect="plain">{{ slot }}</el-tag>
            <el-tag v-for="reward in skill.rewards" :key="reward" type="success" effect="plain">{{ reward }}</el-tag>
          </div>
          <small>{{ skill.owner }} · {{ skill.campus }} · {{ skill.portfolio }}</small>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>需求浏览</h2>
        <NeedCard
          v-for="need in overview.needs"
          :key="need.id"
          :need="need"
          :current-user="overview.profile.name"
          @respond="openRespond"
          @select="handleSelect"
        />
        <RespondDialog v-model="respondDialogVisible" :need="respondingNeed" :submitting="submitting" @submit="submitResponse" />
      </div>

      <div class="panel">
        <h2>智能匹配</h2>
        <FeatureCard v-for="match in overview.matches" :key="match.id" :title="`${match.provider} × ${match.learner}`" :description="match.recommendation">
          <template #tag><el-tag type="warning">{{ match.score }}%</el-tag></template>
          <p class="muted">{{ match.offerSkill }} ↔ {{ match.wantedSkill }}</p>
          <div class="tag-row">
            <el-tag v-for="slot in match.commonSlots" :key="slot">{{ slot }}</el-tag>
          </div>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>预约确认</h2>
        <el-timeline>
          <el-timeline-item v-for="item in overview.appointments" :key="item.id" :timestamp="item.time">
            <strong>{{ item.pair }}</strong>
            <p>
              {{ item.place }}
              <el-tag :type="appointmentTag(item.status)" size="small">{{ item.status }}</el-tag>
            </p>
            <p class="muted">{{ item.agenda }}</p>
          </el-timeline-item>
        </el-timeline>
      </div>

      <div class="panel profile-panel">
        <div>
          <h2>个人主页与技能墙</h2>
          <h3>{{ overview.profile.name }}</h3>
          <p>{{ overview.profile.major }} · {{ overview.profile.creditLevel }}</p>
          <el-progress :percentage="overview.profile.creditScore" />
          <ul>
            <li v-for="item in overview.profile.history" :key="item">{{ item }}</li>
          </ul>
        </div>
        <RadarChart :radar="overview.profile.radar" />
      </div>

      <div class="panel">
        <h2>评价信用</h2>
        <FeatureCard v-for="review in overview.reviews" :key="review.id" :title="`${review.from} → ${review.to}`" :description="review.content">
          <template #tag><el-rate :model-value="review.rating" disabled size="small" /></template>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>消息通知</h2>
        <FeatureCard v-for="conversation in overview.messages" :key="conversation.id" :title="conversation.withUser" :description="conversation.messages.join(' / ')">
          <template #tag><el-badge :value="conversation.unread" /></template>
        </FeatureCard>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import AppHeader from '../components/AppHeader.vue';
import FeatureCard from '../components/FeatureCard.vue';
import MetricCard from '../components/MetricCard.vue';
import RadarChart from '../components/RadarChart.vue';
import NeedCard from '../components/needs/NeedCard.vue';
import RespondDialog from '../components/needs/RespondDialog.vue';
import { fetchOverview } from '../services/storage.service';
import { selectNeedCandidate, submitNeedResponse } from '../services/need.service';
import { APPOINTMENT_STATUS_TAG, type TagType } from '../constants/need.constants';
import type { NeedResponse, NeedView, Overview, ResponsePayload } from '../types/domain';

const overview = ref<Overview | null>(null);
const loading = ref(true);
const error = ref('');
const respondDialogVisible = ref(false);
const respondingNeed = ref<NeedView | null>(null);
const submitting = ref(false);

async function loadOverview() {
  try {
    overview.value = await fetchOverview();
    error.value = '';
  } catch (err) {
    error.value = err instanceof Error ? err.message : '加载失败';
  } finally {
    loading.value = false;
  }
}

onMounted(loadOverview);

function openRespond(need: NeedView) {
  respondingNeed.value = need;
  respondDialogVisible.value = true;
}

async function submitResponse(payload: ResponsePayload) {
  if (!respondingNeed.value) {
    return;
  }
  submitting.value = true;
  try {
    await submitNeedResponse(respondingNeed.value.id, payload.note, payload.slots);
    ElMessage.success('响应已提交，等待发起人选定');
    respondDialogVisible.value = false;
    await loadOverview();
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '响应提交失败');
  } finally {
    submitting.value = false;
  }
}

async function handleSelect(payload: { need: NeedView; candidate: NeedResponse }) {
  const { need, candidate } = payload;
  try {
    await ElMessageBox.confirm(
      `选定 ${candidate.student} 后将按需求时间「${need.expectTime}」生成待双方确认的预约，其余响应停止接收。`,
      '选定候选人',
      { confirmButtonText: '确定选定', cancelButtonText: '再想想', type: 'warning' },
    );
  } catch {
    return;
  }
  try {
    await selectNeedCandidate(need.id, candidate.id);
    ElMessage.success(`已选定 ${candidate.student}，预约待双方确认`);
  } catch (err) {
    // 时段冲突等业务异常：提示原因，需求仍可改选其他人
    ElMessage.error(err instanceof Error ? err.message : '选定失败');
  } finally {
    await loadOverview();
  }
}

function appointmentTag(status: string): TagType {
  return APPOINTMENT_STATUS_TAG[status] ?? 'info';
}
</script>
