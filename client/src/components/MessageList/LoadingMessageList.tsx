import { VFC } from "react";
import { Flex, Stack, SkeletonCircle, SkeletonText } from "@chakra-ui/react";

const LoadingMessageList: VFC = () => {
  return (
    <Stack pb="5">
      <SkeletonText noOfLines={2} w="200px" mb="2" />
      <Stack spacing="24px">
        {Array(5)
          .fill(0)
          .map((v) => (
            <Flex align="center" justify="space-between">
              <SkeletonCircle size="12" />
              <Stack flex="1" ml="4">
                <SkeletonText noOfLines={1} w="200px" mb="2" />
                <SkeletonText noOfLines={3} w="full" />
              </Stack>
            </Flex>
          ))}
      </Stack>
    </Stack>
  );
};

export default LoadingMessageList;
