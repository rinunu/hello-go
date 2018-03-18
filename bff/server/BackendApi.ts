import axiosBase, {AxiosResponse} from "axios";
import {Article} from "~/common/Article";

const axios = axiosBase.create({
    baseURL: 'http://api:8080',
    headers: {
        'ContentType': 'application/json',
    },
    responseType: 'json'
});

export default {
    async get<T>(url: string) {
        const res: AxiosResponse = await axios.get(url);
        return res.data;
    },

    async post<T>(url: string, data: any) {
        const res: AxiosResponse = await axios.post(url, data);
        return res.data;
    }
}
