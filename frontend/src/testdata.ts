const tables = {
    id: 0,
    number: 4,
};

const products = {
    id: 1,
    name: 'Chicken Tenders',
    desc: 'Basic Chicken Tender dinner',
    picture: false,
    price: '12.99',
    catagory: 'chicken',
    wscost: '5.99',
    numOfSides: '2',
};

const catagories = {
    id: 2,
    name: 'generic name',
};

const servers = {
    id: 3,
    name: 'Darrin',
    code: 1234,
};

const orders = {
    id: 4,
    startTime: '11:30pm',
    endTime: '11:35pm',
    table: 4,
    server: 'Darrin',
};

const CustCodes = {
    id: 5,
    startTime: '12:25pm',
    endTime: '12:30pm',
    code: 12,
    order: 1234,
};

const orderItems = {
    id: 6,
    products: 'products',
    order: 123,
};
