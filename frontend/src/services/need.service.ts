import type { Appointment, NeedCard } from '../types/domain';
import { httpGet, httpPost } from './http.service';

export interface RespondPayload {
  offerNote: string;
  freeSlots: string[];
}

export function fetchNeeds(): Promise<NeedCard[]> {
  return httpGet<NeedCard[]>('/needs');
}

export function respondToNeed(needId: number, payload: RespondPayload): Promise<NeedCard> {
  return httpPost<NeedCard>(`/needs/${needId}/responses`, payload);
}

export function selectCandidate(needId: number, responseId: number): Promise<Appointment> {
  return httpPost<Appointment>(`/needs/${needId}/responses/${responseId}/select`);
}

export function confirmAppointment(appointmentId: number): Promise<Appointment> {
  return httpPost<Appointment>(`/appointments/${appointmentId}/confirm`);
}

export function cancelAppointment(appointmentId: number): Promise<Appointment> {
  return httpPost<Appointment>(`/appointments/${appointmentId}/cancel`);
}
