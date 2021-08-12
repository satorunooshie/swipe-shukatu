import { useToast } from "@chakra-ui/react";

type Status = "info" | "warning" | "success" | "error";

export const useShowToast = () => {
  const toast = useToast();
  return (title: string, status: Status = "success") => {
    toast({
      title: title,
      status: status,
      duration: 3000,
      isClosable: true,
    });
  };
};
