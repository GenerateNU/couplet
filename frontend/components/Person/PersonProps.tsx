export type ImageWithCaptionProps = {
  image: string;
  caption?: string;
};

export type PersonHeightProps = {
  feet: number;
  inches: number;
};

export type PersonProps = {
  id: number;
  firstName: string;
  lastName: string;
  age: number;
  pronouns?: string;
  location?: string;
  school?: string;
  work?: string;
  height?: PersonHeightProps;
  bio: string;
  interests: string[];
  habits?: string[];
  instagram: string;
  images: ImageWithCaptionProps[];
};

export type PillProps = {
  textColor: string;
  backgroundColor: string;
  items: string[];
};
