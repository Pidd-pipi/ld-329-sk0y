export type TagType = 'primary' | 'success' | 'info' | 'warning' | 'danger';

/** 可选择的空闲时段，与后端 constants.TimeSlots 保持一致 */
export const TIME_SLOTS = ['周一晚', '周二晚', '周三晚', '周四晚', '周五晚', '周六上午', '周六下午', '周日全天'] as const;

export const NEED_STATUS = { OPEN: '招募中', MATCHED: '已定人' } as const;

export const RESPONSE_STATUS = { PENDING: '待选定', SELECTED: '已入选', REJECTED: '未入选' } as const;

export const APPOINTMENT_STATUS = { PENDING: '待双方确认', CONFIRMED: '双方已确认' } as const;

export const NEED_STATUS_TAG: Record<string, TagType> = {
  [NEED_STATUS.OPEN]: 'success',
  [NEED_STATUS.MATCHED]: 'info',
};

export const RESPONSE_STATUS_TAG: Record<string, TagType> = {
  [RESPONSE_STATUS.PENDING]: 'warning',
  [RESPONSE_STATUS.SELECTED]: 'success',
  [RESPONSE_STATUS.REJECTED]: 'info',
};

export const APPOINTMENT_STATUS_TAG: Record<string, TagType> = {
  [APPOINTMENT_STATUS.PENDING]: 'warning',
  [APPOINTMENT_STATUS.CONFIRMED]: 'success',
};
