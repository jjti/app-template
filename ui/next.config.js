/** @type {import('next').NextConfig} */
const nextConfig = {
  // https://nextjs.org/docs/pages/building-your-application/deploying/static-exports
  output: "export",
  env: {
    SERVER_ADDR: process.env.SERVER_ADDR ?? "",
  },
};

module.exports = nextConfig;
