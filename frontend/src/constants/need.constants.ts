// 响应与预约相关的展示常量集中维护，避免组件内散落魔法字符串。

export const NEED_STATUS_OPEN = 'open';
export const NEED_STATUS_BOOKED = 'booked';

export const RESPONSE_WAITING = 'waiting';
export const RESPONSE_SELECTED = 'selected';
export const RESPONSE_RELEASED = 'released';

export const APPT_PENDING = 'pending';
export const APPT_CONFIRMED = 'confirmed';
export const APPT_CANCELLED = 'cancelled';

export const NEED_STATUS_LABELS: Record<string, string> = {
  open: '招募中',
  booked: '已选定',
};

export const NEED_STATUS_TAG_TYPE: Record<string, 'success' | 'warning' | 'info'> = {
  open: 'success',
  booked: 'warning',
};

export const RESPONSE_STATUS_LABELS: Record<string, string> = {
  waiting: '候选中',
  selected: '已选中',
  released: '未中选',
};

export const RESPONSE_STATUS_TAG_TYPE: Record<string, 'primary' | 'success' | 'info'> = {
  waiting: 'primary',
  selected: 'success',
  released: 'info',
};

export const APPT_STATUS_LABELS: Record<string, string> = {
  pending: '待双方确认',
  confirmed: '双方已确认',
  cancelled: '已取消',
};

export const APPT_STATUS_TAG_TYPE: Record<string, 'warning' | 'success' | 'info'> = {
  pending: 'warning',
  confirmed: 'success',
  cancelled: 'info',
};

export const OFFER_NOTE_MAX_LENGTH = 200;

export const ACTOR_STORAGE_KEY = 'cyskillswap.currentUser';
