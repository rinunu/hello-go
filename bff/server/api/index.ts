import {Router} from 'express'

import articles from './articles';
import sources from './sources';

const router = Router();

router.use(articles);
router.use(sources);

export default router
