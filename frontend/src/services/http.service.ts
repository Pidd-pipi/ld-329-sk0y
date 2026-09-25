import { ACTOR_STORAGE_KEY } from '../constants/need.constants';
import type { ApiErrorBody } from '../types/domain';

const API_BASE = '/api';

function currentActor(): string {
  return localStorage.getItem(ACTOR_STORAGE_KEY) ?? '';
}

function buildHeaders(extra?: HeadersInit): HeadersInit {
  return { ...(extra ?? {}), 'X-User-Name': currentActor() };
}

// ApiException 携带后端错误码与详情，供界面精确展示（如预约冲突）。
export class ApiException extends Error {
  constructor(
    public readonly code: string,
    message: string,
    public readonly status: number,
    public readonly details?: Record<string, unknown> | null,
  ) {
    super(message);
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, {
    ...init,
    headers: buildHeaders(init?.headers),
  });
  if (!response.ok) {
    let body: ApiErrorBody | null = null;
    try {
      body = (await response.json()) as ApiErrorBody;
    } catch {
      body = null;
    }
    throw new ApiException(
      body?.error ?? 'REQUEST_FAILED',
      body?.message ?? '请求失败，请稍后重试',
      response.status,
      body?.details,
    );
  }
  return response.json() as Promise<T>;
}

export function httpGet<T>(path: string): Promise<T> {
  return request<T>(path);
}

export function httpPost<T>(path: string, payload?: unknown): Promise<T> {
  return request<T>(path, {
    method: 'POST',
    headers: payload === undefined ? undefined : { 'Content-Type': 'application/json' },
    body: payload === undefined ? undefined : JSON.stringify(payload),
  });
}
