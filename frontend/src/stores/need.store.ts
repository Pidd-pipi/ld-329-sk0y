import { defineStore } from 'pinia';
import { ElMessage } from 'element-plus';
import {
  ACTOR_STORAGE_KEY,
} from '../constants/need.constants';
import {
  cancelAppointment,
  confirmAppointment,
  respondToNeed,
  selectCandidate,
  type RespondPayload,
} from '../services/need.service';
import { fetchOverview } from '../services/storage.service';
import { ApiException } from '../services/http.service';
import type { Appointment, Overview } from '../types/domain';

const DEFAULT_ACTOR = '林澈';

interface State {
  overview: Overview | null;
  actor: string;
  loading: boolean;
  busyKey: string;
}

export const useNeedStore = defineStore('need', {
  state: (): State => ({
    overview: null,
    actor: localStorage.getItem(ACTOR_STORAGE_KEY) ?? DEFAULT_ACTOR,
    loading: false,
    busyKey: '',
  }),

  getters: {
    slots: (s) => s.overview?.slots ?? [],
  },

  actions: {
    // 切换演示身份后重新拉取（响应/预约状态随当前同学变化）。
    setActor(name: string) {
      this.actor = name;
      localStorage.setItem(ACTOR_STORAGE_KEY, name);
      void this.loadOverview();
    },

    async loadOverview() {
      this.loading = true;
      try {
        this.overview = await fetchOverview();
      } catch (err) {
        ElMessage.error(err instanceof Error ? err.message : '加载看板失败');
      } finally {
        this.loading = false;
      }
    },

    async runAction(key: string, action: () => Promise<unknown>, success: string) {
      if (this.busyKey) return;
      this.busyKey = key;
      try {
        await action();
        await this.loadOverview();
        ElMessage.success(success);
      } catch (err) {
        this.reportError(err);
      } finally {
        this.busyKey = '';
      }
    },

    reportError(err: unknown) {
      if (err instanceof ApiException) {
        ElMessage({ type: 'error', message: err.message, duration: 4200, showClose: true });
      } else {
        ElMessage.error(err instanceof Error ? err.message : '操作失败，请稍后重试');
      }
    },

    submitResponse(needId: number, payload: RespondPayload, done: () => void) {
      return this.runAction(`respond-${needId}`, async () => {
        await respondToNeed(needId, payload);
        done();
      }, '响应已提交，等待发起人选定');
    },

    select(needId: number, responseId: number) {
      return this.runAction(`select-${needId}-${responseId}`, async () => {
        const apt = await selectCandidate(needId, responseId);
        return apt;
      }, '已选定该同学，预约生成，等待双方确认');
    },

    confirm(apt: Appointment) {
      return this.runAction(`confirm-${apt.id}`, () => confirmAppointment(apt.id), '确认成功');
    },

    cancel(apt: Appointment) {
      return this.runAction(`cancel-${apt.id}`, () => cancelAppointment(apt.id), '预约已取消，需求重新开放，可改选其他同学');
    },
  },
});
