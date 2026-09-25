import type { Overview } from '../types/domain';
import { httpGet } from './http.service';

export function fetchOverview(): Promise<Overview> {
  return httpGet<Overview>('/dashboard/overview');
}
