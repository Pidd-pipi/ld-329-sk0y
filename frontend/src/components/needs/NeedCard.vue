<template>
  <FeatureCard :title="need.title" :description="need.description">
    <template #tag>
      <el-tag :type="needTag" size="small">{{ need.status }}</el-tag>
    </template>

    <p class="muted">{{ need.requester }} · {{ need.campus }} · {{ need.expectTime }} · {{ need.budgetType }}</p>
    <div class="tag-row">
      <el-tag size="small" effect="plain">{{ need.category }}</el-tag>
      <el-tag size="small" effect="plain" type="warning">{{ need.responses }} 人响应</el-tag>
    </div>

    <div v-if="need.myResponse" class="need-section">
      <div class="need-section__head">
        <span>我的响应</span>
        <el-tag :type="responseTag(need.myResponse.status)" size="small">{{ need.myResponse.status }}</el-tag>
      </div>
      <p>{{ need.myResponse.note }}</p>
      <div class="tag-row">
        <el-tag v-for="slot in need.myResponse.slots" :key="slot" size="small" effect="plain">{{ slot }}</el-tag>
      </div>
    </div>

    <div v-if="need.myAppointment" class="need-section">
      <div class="need-section__head">
        <span>我的预约</span>
        <el-tag :type="appointmentTag(need.myAppointment.status)" size="small">{{ need.myAppointment.status }}</el-tag>
      </div>
      <p>{{ need.myAppointment.time }} · {{ need.myAppointment.place }}</p>
      <p class="muted">{{ need.myAppointment.agenda }}</p>
    </div>

    <el-button v-if="canRespond" type="primary" size="small" @click="$emit('respond', need)">响应需求</el-button>
    <p v-else-if="closedForMe" class="muted">该需求已定人，停止接收响应</p>

    <CandidateList
      v-if="isMine"
      :candidates="need.candidates ?? []"
      :selectable="need.status === NEED_STATUS.OPEN"
      @select="$emit('select', { need, candidate: $event })"
    />
  </FeatureCard>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import FeatureCard from '../FeatureCard.vue';
import CandidateList from './CandidateList.vue';
import {
  APPOINTMENT_STATUS_TAG,
  NEED_STATUS,
  NEED_STATUS_TAG,
  RESPONSE_STATUS_TAG,
  type TagType,
} from '../../constants/need.constants';
import type { NeedResponse, NeedView } from '../../types/domain';

const props = defineProps<{ need: NeedView; currentUser: string }>();
defineEmits<{
  respond: [need: NeedView];
  select: [payload: { need: NeedView; candidate: NeedResponse }];
}>();

const isMine = computed(() => props.need.requester === props.currentUser);
const canRespond = computed(
  () => !isMine.value && props.need.status === NEED_STATUS.OPEN && !props.need.myResponse,
);
const closedForMe = computed(
  () => !isMine.value && props.need.status === NEED_STATUS.MATCHED && !props.need.myResponse,
);
const needTag = computed<TagType>(() => NEED_STATUS_TAG[props.need.status] ?? 'info');

function responseTag(status: string): TagType {
  return RESPONSE_STATUS_TAG[status] ?? 'info';
}

function appointmentTag(status: string): TagType {
  return APPOINTMENT_STATUS_TAG[status] ?? 'info';
}
</script>
