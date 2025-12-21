import { useParams } from "react-router";

export const Commodity = () => {
  const { commodity } = useParams();

  return <h1>{commodity} page</h1>;
};
