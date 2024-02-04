import { IService } from "./Service";

export type IBudget = {
  id: number;
  price: number;
  date: string;
  userID: string;
  clientID: number;
  clientName: string;
  budgetServices: IService[];
};

export type ICreateBudgetRequest = {
  price: number;
  userID: string;
  clientID: number;
  clientName: string;
  serviceIDs: number[];
};
