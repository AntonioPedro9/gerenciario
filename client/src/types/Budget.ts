import { IJob } from "./Job";

export type IBudget = {
  id: number;
  userID: string;
  customerID: number;
  budgetDate: string;
  scheduledDate?: string;
  budgetJobs: IJob[];
  vehicle: string;
  licensePlate: string;
  price: number;
};

export type ICreateBudgetRequest = {
  userID: string;
  customerID: number;
  jobIDs: number[];
  vehicle: string;
  licensePlate: string;
  price: number;
};

export type IListBudgets = {
  id: number;
  userID: string;
  customerID: number;
  customerName: string;
  customerPhone: string;
  budgetDate: string;
  scheduledDate?: string;
  budgetJobs: IJob[];
  vehicle: string;
  licensePlate: string;
  price: number;
};
