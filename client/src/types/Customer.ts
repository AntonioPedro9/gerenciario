export type ICustomer = {
  id: number;
  cpf: string;
  name: string;
  email: string;
  phone: string;
  userID: string;
};

export type ICreateCustomerRequest = {
  cpf: string;
  name: string;
  email?: string;
  phone: string;
  userID: string;
};

export type IUpdateCustomerRequest = {
  id: number;
  cpf?: string | null;
  name?: string | null;
  email?: string | null;
  phone?: string | null;
  userID: string;
};
