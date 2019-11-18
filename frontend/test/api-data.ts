import {
  Order,
  Product,
  Category,
  CustCode,
  OrderItem
} from "@/graphql/schema";

const categories: Category[] = [
  {
    id: 1,
    name: "Test Category",
    products: []
  }
];

const products: Product[] = [
  {
    id: 1,
    name: "Test Product 1",
    desc: "",
    picture: "",
    price: 999,
    category: categories[0],
    wscost: 523,
    num_of_sides: 0
  }
];

const testOrderTemplate: Order = {
  id: 2,
  table: {
    id: 1,
    num: 1,
    orders: []
  },
  cust_code: null,
  items: [],
  payments: [],
  start_time: "",
  end_time: null,
  server: {
    id: 1,
    name: "Server 1",
    code: 111,
    manager: false
  }
};

const orderItems: OrderItem[] = [
  {
    id: 1,
    order: testOrderTemplate,
    products: [products[0]]
  }
];

const custCodes: CustCode[] = [
  {
    id: 1,
    code: "TEST12",
    order: testOrderTemplate,
    start_time: "",
    end_time: null
  }
];

export function genTestOrder(custCode: boolean = false): Order {
  const testOrder = JSON.parse(JSON.stringify(testOrderTemplate));
  testOrder.table.orders.push(testOrder);
  testOrder.items.push(orderItems[0]);

  if (custCode) {
    testOrder.cust_code = custCodes[0];
  }

  return testOrder;
}
