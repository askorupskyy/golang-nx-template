import { defineConfig } from 'orval';

export default defineConfig({
  petstore: {
    output: {
      mode: 'tags-split',
      target: 'build/api',
      schemas: 'build/api/model',
      client: 'react-query',
      mock: false,
      prettier: true,
      indexFiles: true,
      // TODO: grab this from .env
      baseUrl: 'http://localhost:1323',
    },
    input: {
      target: '../../../router/schema.json',
    },
  },
});
