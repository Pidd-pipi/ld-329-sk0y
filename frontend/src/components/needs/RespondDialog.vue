<template>
  <el-dialog
    :model-value="modelValue"
    title="响应需求"
    width="480px"
    @update:model-value="$emit('update:modelValue', $event)"
    @open="reset"
  >
    <template v-if="need">
      <p class="muted">{{ need.title }} · {{ need.expectTime }}</p>
      <el-form label-position="top">
        <el-form-item label="交换说明" required>
          <el-input
            v-model="note"
            type="textarea"
            :rows="3"
            maxlength="200"
            show-word-limit
            placeholder="说明你能提供什么、希望获得什么回报"
          />
        </el-form-item>
        <el-form-item label="空闲时段" required>
          <el-checkbox-group v-model="slots">
            <el-checkbox v-for="slot in TIME_SLOTS" :key="slot" :value="slot">{{ slot }}</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <el-button @click="$emit('update:modelValue', false)">取消</el-button>
      <el-button type="primary" :loading="submitting" :disabled="!valid" @click="submit">提交响应</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { TIME_SLOTS } from '../../constants/need.constants';
import type { NeedView, ResponsePayload } from '../../types/domain';

const props = defineProps<{ modelValue: boolean; need: NeedView | null; submitting: boolean }>();
const emit = defineEmits<{
  'update:modelValue': [value: boolean];
  submit: [payload: ResponsePayload];
}>();

const note = ref('');
const slots = ref<string[]>([]);

const valid = computed(() => note.value.trim().length > 0 && slots.value.length > 0);

function reset() {
  note.value = '';
  slots.value = [];
}

function submit() {
  if (!valid.value || props.submitting) {
    return;
  }
  emit('submit', { note: note.value.trim(), slots: [...slots.value] });
}
</script>
