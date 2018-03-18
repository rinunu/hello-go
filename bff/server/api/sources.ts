import {Router} from 'express'

import api from "../BackendApi";

const router = Router();

router.get('/sources', async function (req, res, next) {
    const result = await api.get("/rss/sources");
    res.json(result);
});

router.post('/sources', async function (req, res, next) {
    const result = await api.post("/rss/sources", req.body);
    res.json(result);
});

export default router
