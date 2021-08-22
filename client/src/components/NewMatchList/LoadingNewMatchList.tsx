import { VFC } from "react";
import { Flex, Stack, VStack, Skeleton, SkeletonText } from "@chakra-ui/react";

const LoadingNewMatchList: VFC = () => {
  return (
    <Stack mb="6">
      <SkeletonText my="4" noOfLines={2} w="300px" />
      <Flex wrap="nowrap" overflowX="auto">
        {Array(5)
          .fill(0)
          .map((_, i) => (
            <VStack key={i} text="center" display="inline-block" mr="5" minW="100px">
              <Skeleton w="100px" h="150px" />
            </VStack>
          ))}
      </Flex>
    </Stack>
  );
};

export default LoadingNewMatchList;
