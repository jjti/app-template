import { SemanticBlastServiceService } from "../gen";

/**
 * blast queries the api with the sequence.
 */
export const blast = async (seq: string) => {
  const resp = await SemanticBlastServiceService.semanticBlastServiceBlast({
    seq,
  });
};
