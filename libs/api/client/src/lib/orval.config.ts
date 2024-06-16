import { defineConfig } from 'orval';

export default defineConfig({
  petstore: {
    output: {
      mode: 'single',
      target: 'build/api/endpoints/client.ts',
      schemas: 'build/api/model',
      client: 'react-query',
      mock: false,
      prettier: true,
      // TODO: grab this from .env
      baseUrl: 'http://localhost:1323',
    },
    input: {
      target: '../../../router/schema.json',
    },
  },
});
