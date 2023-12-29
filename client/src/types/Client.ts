export type IClient = {
  id: number;
  cpf: string;
  name: string;
  email: string;
  phone: string;
  userID: string;
};

export type ICreateClientRequest = {
  cpf: string;
  name: string;
  email?: string;
  phone: string;
  userID: string;
};

export type IUpdateClientRequest = {
  cpf: string;
  name: string;
  email?: string;
  phone: string;
  userID: string;
};
