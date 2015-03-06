package dynamodb_test

import (
	. "github.com/obieq/goar"
	. "github.com/obieq/goar/db/dynamodb"
	. "github.com/obieq/goar/tests/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dynamodb", func() {
	var (
		ModelS, MK, Sprite, Panamera, Evoque, Bugatti DynamodbAutomobile
		ar                                            *DynamodbAutomobile
	)

	BeforeEach(func() {
		ar = DynamodbAutomobile{}.ToActiveRecord()
	})

	It("should initialize client", func() {
		client := Client()
		Ω(client).ShouldNot(BeNil())
	})

	Context("DB Interactions", func() {
		BeforeEach(func() {
			//ModelS = DynamodbAutomobile{SafetyRating: 5, Automobile: Automobile{Vehicle: Vehicle{Make: "tesla", Year: 2009, Model: "model s"}}}.ToActiveRecord()
			ModelS = DynamodbAutomobile{SafetyRating: 5, Automobile: Automobile{Vehicle: Vehicle{Make: "tesla", Year: 2009, Model: "model s"}}}
			ToAR(&ModelS)
			ModelS.ID = "id1"
			Ω(ModelS.Valid()).Should(BeTrue())

			MK = DynamodbAutomobile{SafetyRating: 3, Automobile: Automobile{Vehicle: Vehicle{Make: "austin healey", Year: 1960, Model: "3000"}}}
			ToAR(&MK)
			MK.ID = "id2"
			Ω(MK.Valid()).Should(BeTrue())

			Sprite = DynamodbAutomobile{SafetyRating: 2, Automobile: Automobile{Vehicle: Vehicle{Make: "austin healey", Year: 1960, Model: "sprite"}}}
			ToAR(&Sprite)
			Sprite.ID = "id3"
			Ω(Sprite.Valid()).Should(BeTrue())

			Panamera = DynamodbAutomobile{SafetyRating: 5, Automobile: Automobile{Vehicle: Vehicle{Make: "porsche", Year: 2010, Model: "panamera"}}}
			ToAR(&Panamera)
			Panamera.ID = "id4"
			Ω(Panamera.Valid()).Should(BeTrue())

			Evoque = DynamodbAutomobile{SafetyRating: 1, Automobile: Automobile{Vehicle: Vehicle{Make: "land rover", Year: 2013, Model: "evoque"}}}
			ToAR(&Evoque)
			Evoque.ID = "id5"
			Ω(Evoque.Valid()).Should(BeTrue())

			Bugatti = DynamodbAutomobile{SafetyRating: 4, Automobile: Automobile{Vehicle: Vehicle{Make: "bugatti", Year: 2013, Model: "veyron"}}}
			ToAR(&Bugatti)
			Bugatti.ID = "id6"
			Ω(Bugatti.Valid()).Should(BeTrue())
		})

		Context("Persistance", func() {
			It("should create a model and find it by id", func() {
				success, err := ModelS.Save()

				Ω(ModelS.ModelName()).Should(Equal("DynamodbAutomobiles"))
				Ω(err).Should(BeNil())
				Ω(success).Should(BeTrue())

				result, err := DynamodbAutomobile{}.ToActiveRecord().Find(ModelS.ID)
				Ω(err).Should(BeNil())
				Ω(result).ShouldNot(BeNil())
				model := result.(*DynamodbAutomobile)
				Ω(model.ID).Should(Equal(ModelS.ID))
				Ω(model.CreatedAt).ShouldNot(BeNil())
			})

			//It("should not create a model using an existing id", func() {
			//Sprite.Delete()
			//Ω(Sprite.Save()).Should(BeTrue())

			//// reset CreatedAt
			//Sprite.CreatedAt = nil
			//success, err := Sprite.Save() // id is still the same, so save should fail
			//Ω(err).To(HaveOccurred())
			//Ω(success).Should(BeFalse())
			//})

			It("should update an existing model", func() {
				Sprite.Delete()
				Ω(Sprite.Save()).Should(BeTrue())
				year := Sprite.Year
				modelName := Sprite.Model

				// create
				result, _ := ar.Find(Sprite.ID)
				Ω(result).ShouldNot(BeNil())
				dbModel := result.(*DynamodbAutomobile).ToActiveRecord()
				Ω(dbModel.ID).Should(Equal(Sprite.ID))
				Ω(dbModel.CreatedAt).ShouldNot(BeNil())
				Ω(dbModel.UpdatedAt).Should(BeNil())

				// update
				dbModel.Year += 1
				dbModel.Model += " updated"
				Ω(dbModel.Save()).Should(BeTrue())

				// verify updates
				result, err := ar.Find(Sprite.ID)
				Expect(err).NotTo(HaveOccurred())
				Ω(result).ShouldNot(BeNil())
				Ω(dbModel.Year).Should(Equal(year + 1))
				Ω(dbModel.Model).Should(Equal(modelName + " updated"))
				Ω(dbModel.CreatedAt).ShouldNot(BeNil())
				Ω(dbModel.UpdatedAt).ShouldNot(BeNil())
			})

			It("should delete an existing model", func() {
				// create and verify
				Ω(MK.Save()).Should(BeTrue())
				result, err := ar.Find(MK.ID)
				Expect(err).NotTo(HaveOccurred())
				Ω(result).ShouldNot(BeNil())
				Ω(MK.ID).Should(Equal(result.(*DynamodbAutomobile).ID))

				// delete
				err = MK.Delete()
				Ω(err).NotTo(HaveOccurred())

				// verify delete
				result, _ = ar.Find(MK.ID)
				Ω(result).Should(BeNil())
			})
		})
	})
})
