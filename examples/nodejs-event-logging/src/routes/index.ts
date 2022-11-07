// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

import { EventListenerList } from '@steadybit/event-kit-api';
import express from 'express';

export const router = express.Router();

router.get('/events', (_, res) => {
	const response: EventListenerList = {
		eventListeners: [
			{
				method: 'get',
				path: '/events/logging'
			}
		]
	};
	res.json(response);
});
