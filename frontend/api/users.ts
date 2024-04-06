import client from "./client";

export default async function getMatchesByUserId(uuid: string) {
  const { data, error } = await client.GET("/matches/{id}", {
    params: {
      path: {
        id: uuid
      }
    }
  });

  if (error) {
    console.error(error);
    throw new Error("Failed to get user matches by ID");
  }

  return data;
}
