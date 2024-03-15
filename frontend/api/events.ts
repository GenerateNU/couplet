import client from "./client";

export default async function getAllEvents() {
  const req = {
    limit: 5,
    offset: 0
  };

  const { data, error } = await client.GET("/events", {
    params: { query: req }
  });

  if (error) {
    console.error(error);
    return [];
  }

  return data;
}
