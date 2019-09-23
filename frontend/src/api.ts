
import axios from 'axios';

export default class ApiClass {

    /*    public getAllDataTest(URL: string, apiData: any) {
           axios
               .get(URL)
               .then((response) => (apiData = response))
               .catch((err) => (err));
       } */

    public getAllTestData(URL: string, apiData: any) {
        /*         this.clicked = !this.clicked;*/
        fetch("https://jsonplaceholder.typicode.com/todos")
            .then((response) => response.json())
            .then((json) => (apiData = json));
    }
}








/*
export interface

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
      fetch(`/api/orders/table/${table}`)
     .then(data =>data.json())
     .then((data: Table) => callback(data));
}
 */

