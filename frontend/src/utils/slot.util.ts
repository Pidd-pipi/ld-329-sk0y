import type { SlotOption } from '../types/domain';

// 时段编码 -> 中文标签的统一转换入口，组件内不直接硬编码时段文案。
export function slotLabel(code: string, slots: SlotOption[]): string {
  return slots.find((s) => s.code === code)?.label ?? code;
}
