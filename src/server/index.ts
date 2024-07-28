/**
 * This a minimal tRPC server
 */
import {createHTTPServer} from '@trpc/server/adapters/standalone';
import {z} from 'zod';
import {db} from './db.js';
import {publicProcedure, router} from './trpc.js';
import {AppRouterClientImpl, ByIdRequest, ByIdResponse} from "../_generated/app.js";

const mockRpc = {
    request: async (service: string, method: string, data: Uint8Array): Promise<Uint8Array> => {
        if (service === 'appRouter.AppRouter' && method === 'ById') {
            // Mock response data
            const response = ByIdResponse.encode({ users: [{ name: 'sachinraja' }] }).finish();
            return Promise.resolve(response);
        }
        return Promise.reject(new Error('Method not implemented'));
    },
};

const appRouterClient = new AppRouterClientImpl(mockRpc);

const appRouter = router({
    byId: publicProcedure.input(z.string()).query(async (opts) => {
        const { input } = opts;
        console.log(input)

        const request = ByIdRequest.create({ id: input });
        console.log(request)

        // Call the ById method
        const response = await appRouterClient.ById(request);

        // Return the decoded response
        return response.users;
    }),
});

export type AppRouter = typeof appRouter;

const server = createHTTPServer({
  router: appRouter,
});

server.listen(3000);
