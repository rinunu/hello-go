import {Router} from 'express'

import api from "../BackendApi";

const router = Router();

router.get('/sources/:id/articles', async function (req, res, next) {
    const id = parseInt(req.params.id);
    const result = await api.get(`/rss/sources/${id}/articles`);
    res.json(result);
});

export default router
