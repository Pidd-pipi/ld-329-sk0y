import type { Appointment, NeedResponse } from '../types/domain';
import { AppException } from '../errors/AppException';

const API_BASE = '/api';

interface ErrorBody {
  code?: string;
  message?: string;
}

/** 后端业务异常统一为 { code, message }，解析失败时给出默认提示 */
async function ensureOk(response: Response, fallback: string): Promise<void> {
  if (response.ok) {
    return;
  }
  let code = 'APP_ERROR';
  let message = fallback;
  try {
    const body = (await response.json()) as ErrorBody;
    if (body.code) code = body.code;
    if (body.message) message = body.message;
  } catch {
    // 保留默认提示
  }
  throw new AppException(message, code);
}

/** 同学提交响应：交换说明 + 空闲时段 */
export async function submitNeedResponse(needId: number, note: string, slots: string[]): Promise<NeedResponse> {
  const response = await fetch(`${API_BASE}/needs/${needId}/responses`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ note, slots }),
  });
  await ensureOk(response, '响应提交失败，请稍后重试');
  return response.json() as Promise<NeedResponse>;
}

/** 发起人选定候选人，生成待双方确认的预约 */
export async function selectNeedCandidate(needId: number, responseId: number): Promise<Appointment> {
  const response = await fetch(`${API_BASE}/needs/${needId}/select`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ responseId }),
  });
  await ensureOk(response, '选定失败，请稍后重试');
  return response.json() as Promise<Appointment>;
}
