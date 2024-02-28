export type IJob = {
  id: number;
  name: string;
  description: string;
  duration: number;
  price: number;
  userID: string;
};

export type ICreateJobRequest = {
  name: string;
  description?: string;
  duration?: number;
  price?: number;
  userID: string;
};

export type IUpdateJobRequest = {
  id: number;
  name: string;
  description?: string;
  duration?: number;
  price?: number;
  userID: string;
};
