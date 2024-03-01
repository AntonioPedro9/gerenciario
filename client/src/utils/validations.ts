import * as yup from "yup";

const customerSchema = yup.object().shape({
  cpf: yup.string(),
  name: yup.string().required("Preencha todos os campos obrigatórios"),
  email: yup.string().email("Email inválido"),
  phone: yup
    .string()
    .matches(/^\d{10,11}$/, "Telefone deve ter 10 ou 11 dígitos")
    .required("Preencha todos os campos obrigatórios"),
});

const jobSchema = yup.object().shape({
  name: yup.string().required("Preencha todos os campos obrigatórios"),
  description: yup.string(),
  duration: yup.number().integer().min(0, "Duração deve ser um número inteiro maior ou igual a 0"),
  price: yup.number().min(0, "O preço deve ser maior ou igual a zero"),
});

const budgetSchema = yup.object().shape({
  vehicle: yup.string().required("Preencha todos os campos obrigatórios"),
  licensePlate: yup.string().required("Preencha todos os campos obrigatórios"),
  price: yup.number().min(0, "O preço deve ser maior ou igual a zero"),
});

export { customerSchema, jobSchema, budgetSchema };
