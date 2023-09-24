import { NextApiRequest, NextApiResponse } from 'next';

export default function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (res.status(200)) {
        console.log("yes");
        res.end('ok')
    }
}