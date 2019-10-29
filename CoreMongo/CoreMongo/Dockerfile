FROM mcr.microsoft.com/dotnet/core/aspnet:3.0-buster-slim AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/core/sdk:3.0-buster AS build
WORKDIR /src
COPY ["CoreMongo/CoreMongo.csproj", "CoreMongo/"]
RUN dotnet restore "CoreMongo/CoreMongo.csproj"
COPY . .
WORKDIR "/src/CoreMongo"
RUN dotnet build "CoreMongo.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "CoreMongo.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "CoreMongo.dll"]