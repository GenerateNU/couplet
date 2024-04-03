import client from "./client";

export default async function getOrgById(uuid: string) {
  const { data, error } = await client.GET("/orgs/{id}", {
    params: {
      path: {
        id: uuid
      }
    }
  });

  if (error) {
    console.error(error);
    throw new Error("Failed to get org by ID");
  }

  return data;
}
