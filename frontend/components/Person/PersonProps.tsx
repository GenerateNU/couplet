import type { components } from "../../api/schema";

type Event = components["schemas"]["Event"];

export type ImageWithCaptionProps = {
  image: string;
  caption?: string;
};

export type PersonHeightProps = {
  feet: number;
  inches: number;
};

export type PersonProps = {
  // Person info
  id: number;
  firstName: string;
  lastName: string;
  age: number;
  pronouns?: string;

  // Basic info
  location?: string;
  school?: string;
  work?: string;
  height?: PersonHeightProps;

  bio: string;
  interests: string[];
  
  // Habits
  relationshipType?: string; 
  religion?: string; 
  politicalAffiliation?: string;
  alchoholFrequency?: string;
  smokingFrequency?: string;
  drugFrequency?: string;
  marijuanaFrequency?: string;

  instagram: string;
  mutualEvents: Event[]; 
  images: ImageWithCaptionProps[];
  isMatched: boolean; 
};

export type PillProps = {
  textColor: string;
  backgroundColor: string;
  items: string[];
};
