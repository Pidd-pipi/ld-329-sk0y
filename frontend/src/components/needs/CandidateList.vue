<template>
  <div class="candidate-list">
    <p class="muted">候选人（{{ candidates.length }}）</p>
    <el-empty v-if="candidates.length === 0" description="暂无响应" :image-size="48" />
    <div v-for="candidate in candidates" :key="candidate.id" class="candidate-item">
      <div class="candidate-item__body">
        <div class="candidate-item__head">
          <strong>{{ candidate.student }}</strong>
          <el-tag :type="statusTag(candidate.status)" size="small">{{ candidate.status }}</el-tag>
        </div>
        <p>{{ candidate.note }}</p>
        <div class="tag-row">
          <el-tag v-for="slot in candidate.slots" :key="slot" size="small" effect="plain">{{ slot }}</el-tag>
        </div>
      </div>
      <el-button
        v-if="selectable && candidate.status === RESPONSE_STATUS.PENDING"
        type="primary"
        size="small"
        @click="$emit('select', candidate)"
      >
        选定
      </el-button>
    </div>
    <p v-if="!selectable" class="muted">需求已定人，停止接收新的响应</p>
  </div>
</template>

<script setup lang="ts">
import { RESPONSE_STATUS, RESPONSE_STATUS_TAG, type TagType } from '../../constants/need.constants';
import type { NeedResponse } from '../../types/domain';

defineProps<{ candidates: NeedResponse[]; selectable: boolean }>();
defineEmits<{ select: [candidate: NeedResponse] }>();

function statusTag(status: string): TagType {
  return RESPONSE_STATUS_TAG[status] ?? 'info';
}
</script>
