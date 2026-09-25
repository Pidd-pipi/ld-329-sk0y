<template>
  <div v-if="appointment" class="appt-banner" :class="`appt-banner--${appointment.status}`">
    <div class="appt-banner__head">
      <el-tag :type="statusTagType" size="small">{{ statusLabel }}</el-tag>
      <span class="appt-banner__pair">{{ appointment.requester }} ↔ {{ appointment.respondent }}</span>
      <span class="muted">{{ appointment.time }} · {{ appointment.place }}</span>
    </div>
    <p class="appt-banner__agenda muted">{{ appointment.agenda }}</p>

    <div v-if="appointment.status !== 'cancelled'" class="appt-banner__confirm">
      <span>
        发起人确认：
        <el-tag size="small" :type="appointment.requesterConfirmed ? 'success' : 'info'">
          {{ appointment.requesterConfirmed ? '已确认' : '待确认' }}
        </el-tag>
      </span>
      <span>
        被选同学确认：
        <el-tag size="small" :type="appointment.respondentConfirmed ? 'success' : 'info'">
          {{ appointment.respondentConfirmed ? '已确认' : '待确认' }}
        </el-tag>
      </span>
    </div>

    <div v-if="isParty && appointment.status === 'pending'" class="appt-banner__actions">
      <el-button
        size="small"
        type="primary"
        :loading="busy"
        :disabled="alreadyConfirmed"
        @click="emit('confirm')"
      >
        {{ alreadyConfirmed ? '已确认，等待对方' : '确认预约' }}
      </el-button>
      <el-button size="small" :loading="busy" @click="onCancel">取消并改选</el-button>
    </div>
    <p v-else-if="appointment.status === 'confirmed'" class="appt-banner__done muted">
      双方已确认，交换安排生效。
    </p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { ElMessageBox } from 'element-plus';
import { APPT_STATUS_LABELS, APPT_STATUS_TAG_TYPE } from '../../constants/need.constants';
import type { Appointment } from '../../types/domain';

const props = defineProps<{
  appointment: Appointment;
  currentUser: string;
  busy: boolean;
}>();
const emit = defineEmits<{
  (e: 'confirm'): void;
  (e: 'cancel'): void;
}>();

const statusLabel = computed(() => APPT_STATUS_LABELS[props.appointment.status] ?? props.appointment.status);
const statusTagType = computed(() => APPT_STATUS_TAG_TYPE[props.appointment.status] ?? 'info');

const isParty = computed(
  () => props.currentUser === props.appointment.requester || props.currentUser === props.appointment.respondent,
);
const alreadyConfirmed = computed(() => {
  if (props.currentUser === props.appointment.requester) return props.appointment.requesterConfirmed;
  if (props.currentUser === props.appointment.respondent) return props.appointment.respondentConfirmed;
  return false;
});

async function onCancel() {
  try {
    await ElMessageBox.confirm(
      '取消后该需求将重新开放，其他候选人恢复候选，确定取消并改选吗？',
      '取消预约',
      { type: 'warning', confirmButtonText: '取消预约并改选', cancelButtonText: '再想想' },
    );
  } catch {
    return;
  }
  emit('cancel');
}
</script>

<style scoped>
.appt-banner { margin-top: 12px; padding: 12px; border-radius: 6px; border: 1px solid #e5e7eb; background: #fafafa; }
.appt-banner--pending { border-color: #f59e0b; background: #fffbeb; }
.appt-banner--confirmed { border-color: #10b981; background: #ecfdf5; }
.appt-banner__head { display: flex; align-items: center; gap: 10px; flex-wrap: wrap; }
.appt-banner__pair { font-weight: 600; }
.appt-banner__agenda { margin: 8px 0; }
.appt-banner__confirm { display: flex; gap: 18px; flex-wrap: wrap; font-size: 13px; }
.appt-banner__actions { margin-top: 10px; display: flex; gap: 8px; }
.appt-banner__done { margin: 8px 0 0; }
</style>
