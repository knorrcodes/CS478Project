// export interface
export interface Table {
    id: number;
    number: number;
};

const tables = [
    {
        id: 5,
        number: 4,
    }
];

export function getCurrentTableOrder(table: number, callback: (data: Table) => void) {
    callback(tables[table]);
    /*  fetch(`/api/orders/table/${table}`)
     .then(data =>data.json())
     .then((data: Table) => callback(data)); */

}
