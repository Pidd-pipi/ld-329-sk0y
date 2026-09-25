<template>
  <main class="page-shell" v-loading="store.loading">
    <AppHeader
      :unread="store.overview?.metrics.unread ?? 0"
      :current-user="store.actor"
      :users="store.overview?.switchableUsers ?? [store.actor]"
      @change-actor="store.setActor($event)"
    />

    <section v-if="store.overview" class="metrics-grid">
      <MetricCard label="已发布技能" :value="store.overview.metrics.skills" />
      <MetricCard label="活跃需求" :value="store.overview.metrics.needs" />
      <MetricCard label="智能匹配" :value="store.overview.metrics.matches" />
      <MetricCard label="生效预约" :value="store.overview.metrics.appointments" />
    </section>

    <el-alert v-if="error" :title="error" type="error" show-icon />

    <section v-if="store.overview" class="workspace-grid">
      <div class="panel">
        <h2>技能发布</h2>
        <FeatureCard v-for="skill in store.overview.skills" :key="skill.id" :title="skill.title" :description="skill.description">
          <template #tag><el-tag>{{ skill.category }} {{ skill.level }}%</el-tag></template>
          <div class="tag-row">
            <el-tag v-for="slot in skill.timeSlots" :key="slot" effect="plain">{{ slot }}</el-tag>
            <el-tag v-for="reward in skill.rewards" :key="reward" type="success" effect="plain">{{ reward }}</el-tag>
          </div>
          <small>{{ skill.owner }} · {{ skill.campus }} · {{ skill.portfolio }}</small>
        </FeatureCard>
      </div>

      <div class="panel panel--wide">
        <NeedBoard />
      </div>

      <div class="panel">
        <h2>智能匹配</h2>
        <FeatureCard v-for="match in store.overview.matches" :key="match.id" :title="`${match.provider} × ${match.learner}`" :description="match.recommendation">
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
          <el-timeline-item
            v-for="item in store.overview.appointments"
            :key="item.id"
            :timestamp="`${item.time} · ${item.place}`"
            :type="timelineType(item.status)"
          >
            <div class="appt-line">
              <strong>{{ item.requester }} ↔ {{ item.respondent }}</strong>
              <el-tag size="small" :type="apptTagType(item.status)">{{ apptStatusLabel(item.status) }}</el-tag>
            </div>
            <p class="muted">
              发起人{{ item.requesterConfirmed ? '✓' : '…' }} / 被选同学{{ item.respondentConfirmed ? '✓' : '…' }}
            </p>
            <p class="muted">{{ item.agenda }}</p>
          </el-timeline-item>
        </el-timeline>
      </div>

      <div class="panel profile-panel">
        <div>
          <h2>个人主页与技能墙</h2>
          <h3>{{ store.overview.profile.name }}</h3>
          <p>{{ store.overview.profile.major }} · {{ store.overview.profile.creditLevel }}</p>
          <el-progress :percentage="store.overview.profile.creditScore" />
          <ul>
            <li v-for="item in store.overview.profile.history" :key="item">{{ item }}</li>
          </ul>
        </div>
        <RadarChart :radar="store.overview.profile.radar" />
      </div>

      <div class="panel">
        <h2>评价信用</h2>
        <FeatureCard v-for="review in store.overview.reviews" :key="review.id" :title="`${review.from} → ${review.to}`" :description="review.content">
          <template #tag><el-rate :model-value="review.rating" disabled size="small" /></template>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>消息通知</h2>
        <FeatureCard v-for="conversation in store.overview.messages" :key="conversation.id" :title="conversation.withUser" :description="conversation.messages.join(' / ')">
          <template #tag><el-badge :value="conversation.unread" /></template>
        </FeatureCard>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import AppHeader from '../components/AppHeader.vue';
import FeatureCard from '../components/FeatureCard.vue';
import MetricCard from '../components/MetricCard.vue';
import RadarChart from '../components/RadarChart.vue';
import NeedBoard from '../components/need/NeedBoard.vue';
import { APPT_STATUS_LABELS, APPT_STATUS_TAG_TYPE } from '../constants/need.constants';
import { useNeedStore } from '../stores/need.store';

const store = useNeedStore();
const error = ref('');

onMounted(async () => {
  try {
    await store.loadOverview();
  } catch (err) {
    error.value = err instanceof Error ? err.message : '加载失败';
  }
});

function apptStatusLabel(status: string): string {
  return APPT_STATUS_LABELS[status] ?? status;
}
function apptTagType(status: string): 'warning' | 'success' | 'info' {
  return APPT_STATUS_TAG_TYPE[status] ?? 'info';
}
function timelineType(status: string): 'warning' | 'success' | 'info' {
  return apptTagType(status);
}
</script>

<style scoped>
.appt-line { display: flex; align-items: center; gap: 8px; }
</style>
