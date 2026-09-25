<template>
  <el-dialog v-model="visible" title="响应需求" width="520px" @closed="reset">
    <el-form label-position="top" @submit.prevent>
      <el-form-item label="交换说明" required>
        <el-input
          v-model="form.offerNote"
          type="textarea"
          :rows="4"
          maxlength="200"
          show-word-limit
          placeholder="介绍你能提供的帮助、相关经验和交换方式……"
        />
      </el-form-item>
      <el-form-item label="我的空闲时段（可多选）" required>
        <el-checkbox-group v-model="form.freeSlots">
          <el-checkbox v-for="slot in slots" :key="slot.code" :value="slot.code">
            {{ slot.label }}
          </el-checkbox>
        </el-checkbox-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="submit">提交响应</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { ElMessage } from 'element-plus';
import { OFFER_NOTE_MAX_LENGTH } from '../../constants/need.constants';
import type { SlotOption } from '../../types/domain';

const props = defineProps<{
  modelValue: boolean;
  slots: SlotOption[];
  submitting: boolean;
}>();
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void;
  (e: 'submit', payload: { offerNote: string; freeSlots: string[] }): void;
}>();

const visible = ref(props.modelValue);
watch(() => props.modelValue, (v) => { visible.value = v; });
watch(visible, (v) => emit('update:modelValue', v));

const form = reactive({ offerNote: '', freeSlots: [] as string[] });

function reset() {
  form.offerNote = '';
  form.freeSlots = [];
}

function submit() {
  const note = form.offerNote.trim();
  if (!note) {
    ElMessage.warning('请填写交换说明');
    return;
  }
  if (note.length > OFFER_NOTE_MAX_LENGTH) {
    ElMessage.warning('交换说明不能超过 200 字');
    return;
  }
  if (form.freeSlots.length === 0) {
    ElMessage.warning('请至少选择一个空闲时段');
    return;
  }
  emit('submit', { offerNote: note, freeSlots: [...form.freeSlots] });
}
</script>
