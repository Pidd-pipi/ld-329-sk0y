<template>
  <article class="need-card">
    <header class="need-card__head">
      <div class="need-card__title">
        <strong>{{ need.title }}</strong>
        <el-tag size="small" :type="needStatusTagType">{{ needStatusLabel }}</el-tag>
      </div>
      <span class="muted">{{ need.category }} · {{ need.campus }} · {{ need.budgetType }}</span>
    </header>

    <p class="need-card__desc">{{ need.description }}</p>

    <div class="tag-row">
      <el-tag type="warning" size="small">期望：{{ need.expectTime }}</el-tag>
      <el-tag size="small" effect="plain">{{ need.responseCount }} 人响应</el-tag>
      <el-tag size="small" effect="plain">发起人：{{ need.requester }}</el-tag>
    </div>

    <!-- 当前用户自己的响应状态 -->
    <el-alert
      v-if="need.myResponse"
      :title="`你已响应：${myResponseStatusText}`"
      :type="myResponseAlertType"
      :closable="false"
      show-icon
      class="need-card__my"
    >
      <div class="my-response">
        <p class="muted">{{ need.myResponse.offerNote }}</p>
        <div class="tag-row">
          <el-tag
            v-for="code in need.myResponse.freeSlots"
            :key="code"
            size="small"
            effect="plain"
          >{{ slotLabel(code, slots) }}</el-tag>
        </div>
      </div>
    </el-alert>

    <!-- 预约状态（待确认 / 已确认 / 已取消改选） -->
    <AppointmentBanner
      v-if="need.appointment"
      :appointment="need.appointment"
      :current-user="currentUser"
      :busy="busyKey.startsWith(`confirm-${need.appointment.id}`) || busyKey.startsWith(`cancel-${need.appointment.id}`)"
      @confirm="emit('confirm', need.appointment!)"
      @cancel="emit('cancel', need.appointment!)"
    />

    <!-- 候选人列表（发起人可见全部，其他人仅见数量） -->
    <CandidateList
      v-if="isRequester"
      :need="need"
      :responses="need.responses"
      :slots="slots"
      :current-user="currentUser"
      :busy-key="busyKey"
      @select="emit('select', $event)"
    />

    <footer class="need-card__foot">
      <el-button
        v-if="canRespond"
        type="primary"
        size="small"
        @click="emit('respond-click')"
      >我要响应</el-button>
      <span v-else-if="isRequester && need.status === 'open'" class="muted">
        你是发起人，可从候选人中选定一人
      </span>
      <span v-else-if="isRequester && need.status === 'booked'" class="muted">
        已选定同学，等待预约双方确认
      </span>
      <span v-else-if="need.myResponse && need.status === 'open'" class="muted">
        已报名，同一条需求无需重复提交
      </span>
      <span v-else-if="!isRequester && need.status === 'booked'" class="muted">
        该需求已选定同学，停止接收响应
      </span>
    </footer>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { AlertProps } from 'element-plus';
import {
  NEED_STATUS_BOOKED,
  NEED_STATUS_LABELS,
  NEED_STATUS_OPEN,
  NEED_STATUS_TAG_TYPE,
  RESPONSE_SELECTED,
  RESPONSE_WAITING,
} from '../../constants/need.constants';
import type { Appointment, NeedCard as NeedCardType, Response, SlotOption } from '../../types/domain';
import { slotLabel } from '../../utils/slot.util';
import AppointmentBanner from './AppointmentBanner.vue';
import CandidateList from './CandidateList.vue';

const props = defineProps<{
  need: NeedCardType;
  slots: SlotOption[];
  currentUser: string;
  busyKey: string;
}>();
const emit = defineEmits<{
  (e: 'respond-click'): void;
  (e: 'select', resp: Response): void;
  (e: 'confirm', apt: Appointment): void;
  (e: 'cancel', apt: Appointment): void;
}>();

const isRequester = computed(() => props.need.requester === props.currentUser);
const needStatusLabel = computed(
  () => NEED_STATUS_LABELS[props.need.status] ?? props.need.status,
);
const needStatusTagType = computed(
  () => NEED_STATUS_TAG_TYPE[props.need.status] ?? 'info',
);

const RELEASED_TEXT = '当前需求已截止';

const canRespond = computed(
  () =>
    !isRequester.value &&
    props.need.status === NEED_STATUS_OPEN &&
    !props.need.appointment &&
    !props.need.myResponse,
);

// 候选人或被选中人看到自己响应的当前含义
const myResponseStatusText = computed(() => {
  const mine = props.need.myResponse;
  if (!mine) return '';
  if (mine.status === RESPONSE_SELECTED) return '发起人已选定你，请确认预约';
  if (mine.status === 'released') return '发起人已选定其他同学，本次未中选';
  if (mine.status === RESPONSE_WAITING) return '等待发起人选定';
  return RELEASED_TEXT;
});

const myResponseAlertType = computed<AlertProps['type']>(() => {
  const mine = props.need.myResponse;
  if (mine?.status === RESPONSE_SELECTED) return 'success';
  if (props.need.status === NEED_STATUS_BOOKED) return 'info';
  return 'warning';
});
</script>

<style scoped>
.need-card { background: #fff; border: 1px solid #e5e7eb; border-radius: 8px; padding: 16px; display: flex; flex-direction: column; gap: 4px; }
.need-card__head { display: flex; align-items: center; justify-content: space-between; gap: 12px; flex-wrap: wrap; }
.need-card__title { display: flex; align-items: center; gap: 8px; }
.need-card__desc { margin: 8px 0; line-height: 1.6; color: #344054; }
.need-card__my { margin: 8px 0; }
.my-response p { margin: 4px 0; }
.need-card__foot { margin-top: 8px; min-height: 28px; display: flex; align-items: center; gap: 10px; flex-wrap: wrap; }
</style>
