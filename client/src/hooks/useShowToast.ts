import { useToast } from "@chakra-ui/react";

type Status = "info" | "warning" | "success" | "error";

export const useShowToast = (title: string, status: Status = "success") => {
  const toast = useToast();
  return () => {
    toast({
      title: title,
      status: status,
      duration: 3000,
      isClosable: true,
    });
  };
};
