export enum Role {
  MANAGER = "MANAGER",
  SERVER = "SERVER",
  CUSTOMER = "CUSTOMER"
}

export interface Product {
  id: number;
  name: string;
  desc: string;
  picture: string;
  price: number;
  category: Category;
  wscost: number;
  num_of_sides: number;
}

export interface Category {
  id: number;
  name: string;
  products: Product[];
}

export interface Server {
  id: number;
  name: string;
  code: number;
  manager: boolean;
}

export interface Order {
  id: number;
  start_time: string;
  end_time: string | null;
  table: Table;
  server: Server;
  items: OrderItem[];
  payments: Payment[];
  cust_code: CustCode | null;
}

export interface Table {
  id: number;
  num: number;
  orders: Order[];
}

export interface CustCode {
  id: number;
  start_time: string;
  end_time: string | null;
  code: string;
  order: Order;
}

export interface OrderItem {
  id: number;
  products: Product[];
  order: Order;
}

export interface Payment {
  id: number;
  order: Order;
  amount: number;
  timestamp: string;
}

export enum OrderStatus {
  ANY = "ANY",
  OPENED = "OPENED",
  CLOSED = "CLOSED"
}

export interface NewProductInput {
  name: string;
  desc: string;
  picture: string;
  price: number;
  category: number;
  wscost: number;
  num_of_sides: number;
}

export interface NewOrderInput {
  table: number;
}

export interface AddPaymentInput {
  order: number;
  amount: number;
}
