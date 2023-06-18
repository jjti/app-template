"use client";

import { blast } from "./api";
import { ChakraProvider } from "@chakra-ui/react";
import { Textarea, Container, Heading, Button, VStack } from "@chakra-ui/react";
import { useState } from "react";

export default function Home() {
  const [seqValue, setSeqValue] = useState("");

  return (
    <ChakraProvider>
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <Container>
          <VStack spacing={2} placeItems={"flex-start"}>
            <Heading>Semantic BLAST</Heading>
            <Textarea
              value={seqValue}
              onChange={(event) => setSeqValue(event.currentTarget.value)}
              size="sm"
              height={300}
            />
            <Button
              onClick={() => {
                blast(seqValue).then((response) => {
                  console.log(response);
                });
              }}
            >
              Submit
            </Button>
          </VStack>
        </Container>
      </main>
    </ChakraProvider>
  );
}
