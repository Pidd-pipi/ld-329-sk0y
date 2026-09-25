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

export interface Need {
  id: number;
  requester: string;
  title: string;
  category: string;
  campus: string;
  expectTime: string;
  slot: string;
  budgetType: string;
  description: string;
  status: string;
  responses: number;
}

export interface NeedResponse {
  id: number;
  needId: number;
  student: string;
  note: string;
  slots: string[];
  status: string;
}

export interface Appointment {
  id: number;
  needId: number;
  pair: string;
  requester: string;
  provider: string;
  time: string;
  slot: string;
  place: string;
  status: string;
  agenda: string;
}

/** 需求卡片视图：附带当前用户的响应、预约与候选人 */
export interface NeedView extends Need {
  myResponse?: NeedResponse;
  myAppointment?: Appointment;
  candidates?: NeedResponse[];
}

/** 提交响应的载荷：交换说明 + 空闲时段 */
export interface ResponsePayload {
  note: string;
  slots: string[];
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
  categories: string[];
  metrics: Record<string, number>;
  skills: Skill[];
  needs: NeedView[];
  matches: Match[];
  appointments: Appointment[];
  reviews: Review[];
  messages: Conversation[];
  profile: Profile;
}
