echo "Make sure you have an account at hub.docker.com"
echo "Then you should login by invoking: docker login -u <your-user-name> --password-stdin"
echo "Your prefix should be: docker.io/<your-user-name>"
echo "Press ctrl+c if you miss any of the prerequisites."

DEFAULT_PREFIX="docker.io/gofrendi"
read -p "Put your docker prefix (default: ${DEFAULT_PREFIX}): " PREFIX
if [ -z "${PREFIX}" ]
then
    PREFIX="${DEFAULT_PREFIX}"
fi

docker build -t ${PREFIX}/alta-batch-3-go-app .
docker image push ${PREFIX}/alta-batch-3-go-app:latest