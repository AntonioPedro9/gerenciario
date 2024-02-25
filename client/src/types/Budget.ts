import { IService } from "./Service";

export type IBudget = {
  id: number;
  price: number;
  vehicle: string;
  licensePlate: string;
  date: string;
  userID: string;
  clientID: number;
  clientName: string;
  budgetServices: IService[];
};

export type ICreateBudgetRequest = {
  price: number;
  vehicle: string;
  licensePlate: string;
  userID: string;
  clientID: number;
  clientName: string;
  serviceIDs: number[];
};
