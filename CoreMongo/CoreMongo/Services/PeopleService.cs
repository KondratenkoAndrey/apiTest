using CoreMongo.Models;
using MongoDB.Driver;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CoreMongo.Services
{
    public class PeopleService
    {
        private readonly IMongoCollection<Person> _people;

        public PeopleService(IPeopleDatabaseSettings settings)
        {
            var client = new MongoClient(settings.ConnectionString);
            var database = client.GetDatabase(settings.DatabaseName);

            _people = database.GetCollection<Person>(settings.PeopleCollectionName);
        }

        public Task<List<Person>> Get() =>
            _people.Find(person => true).ToListAsync();

        public Person Get(string id) =>
            _people.Find<Person>(person => person.Id == id).FirstOrDefault();

        public Person Create(Person person)
        {
            _people.InsertOne(person);
            return person;
        }

        public void Update(string id, Person personIn) =>
            _people.ReplaceOne(person => person.Id == id, personIn);

        public void Remove(Person personIn) =>
            _people.DeleteOne(person => person.Id == personIn.Id);

        public void Remove(string id) =>
            _people.DeleteOne(person => person.Id == id);
    }
}
