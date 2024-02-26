import { IService } from "./Service";

export type IBudget = {
  id: number;
  userID: string;
  clientID: number;
  clientName: string;
  clientPhone: string;
  date: string;
  budgetServices: IService[];
  vehicle: string;
  licensePlate: string;
  price: number;
};

export type ICreateBudgetRequest = {
  userID: string;
  clientID: number;
  clientName: string;
  clientPhone: string;
  serviceIDs: number[];
  vehicle: string;
  licensePlate: string;
  price: number;
};
