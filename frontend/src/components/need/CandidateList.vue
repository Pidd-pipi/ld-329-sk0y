<template>
  <div class="candidate-list">
    <p class="candidate-list__title">
      候选人（{{ responses.length }}）
      <span v-if="need.status === 'booked'" class="muted">· 响应已截止</span>
    </p>
    <el-scrollbar max-height="260px">
      <div v-for="resp in responses" :key="resp.id" class="candidate">
        <div class="candidate__head">
          <strong>{{ resp.respondent }}</strong>
          <el-tag size="small" :type="responseTagType(resp.status)">
            {{ responseStatusLabel(resp.status) }}
          </el-tag>
          <el-button
            v-if="canSelect"
            size="small"
            type="primary"
            plain
            class="candidate__select"
            :loading="busyKey === `select-${need.id}-${resp.id}`"
            @click="emit('select', resp)"
          >
            选定 TA
          </el-button>
        </div>
        <p class="candidate__note">{{ resp.offerNote }}</p>
        <div class="tag-row">
          <el-tag
            v-for="code in resp.freeSlots"
            :key="code"
            size="small"
            :effect="code === need.slotCode ? 'dark' : 'plain'"
            :type="code === need.slotCode ? 'success' : 'info'"
          >
            {{ slotLabel(code, slots) }}
          </el-tag>
        </div>
      </div>
      <el-empty v-if="responses.length === 0" description="还没有人响应" :image-size="60" />
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import {
  NEED_STATUS_OPEN,
  RESPONSE_STATUS_LABELS,
  RESPONSE_STATUS_TAG_TYPE,
} from '../../constants/need.constants';
import type { NeedCard, Response, SlotOption } from '../../types/domain';
import { slotLabel } from '../../utils/slot.util';

const props = defineProps<{
  need: NeedCard;
  responses: Response[];
  slots: SlotOption[];
  currentUser: string;
  busyKey: string;
}>();
const emit = defineEmits<{
  (e: 'select', resp: Response): void;
}>();

const canSelect = computed(
  () =>
    props.need.status === NEED_STATUS_OPEN &&
    props.need.requester === props.currentUser &&
    !props.need.appointment,
);

function responseStatusLabel(status: string): string {
  return RESPONSE_STATUS_LABELS[status] ?? status;
}
function responseTagType(status: string): 'primary' | 'success' | 'info' {
  return RESPONSE_STATUS_TAG_TYPE[status] ?? 'info';
}
</script>

<style scoped>
.candidate-list__title { margin: 10px 0 8px; font-weight: 600; }
.candidate { padding: 10px; border: 1px solid #e5e7eb; border-radius: 6px; margin-bottom: 8px; background: #fff; }
.candidate__head { display: flex; align-items: center; gap: 8px; }
.candidate__select { margin-left: auto; }
.candidate__note { margin: 6px 0; font-size: 13px; line-height: 1.6; color: #475467; }
</style>
