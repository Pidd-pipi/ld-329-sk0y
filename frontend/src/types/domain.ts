export interface Skill {
  id: number;
  owner: string;
  title: string;
  category: string;
  level: number;
  campus: string;
  description: string;
  timeSlots: string[];
  rewards: string[];
  portfolio: string;
}

// 需求状态
export const NEED_STATUS_OPEN = 'open';
export const NEED_STATUS_BOOKED = 'booked';

// 响应状态
export const RESPONSE_WAITING = 'waiting';
export const RESPONSE_SELECTED = 'selected';
export const RESPONSE_RELEASED = 'released';

// 预约状态
export const APPT_PENDING = 'pending';
export const APPT_CONFIRMED = 'confirmed';
export const APPT_CANCELLED = 'cancelled';

export interface Need {
  id: number;
  requester: string;
  title: string;
  category: string;
  campus: string;
  expectTime: string;
  slotCode: string;
  budgetType: string;
  description: string;
  status: string;
}

export interface Response {
  id: number;
  needId: number;
  respondent: string;
  offerNote: string;
  freeSlots: string[];
  status: string;
  createdAt: string;
}

export interface Appointment {
  id: number;
  needId: number;
  requester: string;
  respondent: string;
  slotCode: string;
  time: string;
  place: string;
  agenda: string;
  status: string;
  requesterConfirmed: boolean;
  respondentConfirmed: boolean;
  createdAt: string;
}

export interface NeedCard extends Need {
  responses: Response[];
  responseCount: number;
  myResponse?: Response;
  appointment?: Appointment;
}

export interface SlotOption {
  code: string;
  label: string;
}

export interface Match {
  id: number;
  provider: string;
  learner: string;
  offerSkill: string;
  wantedSkill: string;
  score: number;
  commonSlots: string[];
  recommendation: string;
}

export interface Review {
  id: number;
  from: string;
  to: string;
  rating: number;
  content: string;
}

export interface Conversation {
  id: number;
  withUser: string;
  unread: number;
  messages: string[];
}

export interface Profile {
  name: string;
  major: string;
  creditScore: number;
  creditLevel: string;
  skillWall: Skill[];
  radar: Record<string, number>;
  history: string[];
  reviews: Review[];
}

export interface Overview {
  service: string;
  currentUser: string;
  switchableUsers: string[];
  categories: string[];
  slots: SlotOption[];
  metrics: Record<string, number>;
  skills: Skill[];
  needs: NeedCard[];
  matches: Match[];
  appointments: Appointment[];
  reviews: Review[];
  messages: Conversation[];
  profile: Profile;
}

// 后端 4xx 响应体
export interface ApiErrorBody {
  error: string;
  message: string;
  details?: Record<string, unknown> | null;
}
