package mssql

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	. "github.com/obieq/goar"
)

type ArMsSql struct {
	ActiveRecord
	ID string `json:"id,omitempty"`
	Timestamps
}

var (
	client *sql.DB
)

var connectOpts = func() map[string]string {
	opts := make(map[string]string)
	if envs, err := godotenv.Read(); err != nil {
		log.Fatal("Error loading mssql .env file")
	} else {
		log.Println("OBIE:", envs)
		opts["server"] = envs["MSSQL_SERVER"]
		opts["port"] = envs["MSSQL_PORT"]
		opts["username"] = envs["MSSQL_USERNAME"]
		opts["password"] = envs["MSSQL_PASSWORD"]
		opts["debug"] = envs["MSSQL_DEBUG"]
	}

	return opts
}

func connect() (client *sql.DB) {
	opts := connectOpts()
	server := opts["server"]
	username := opts["username"]
	password := opts["password"]
	port, err := strconv.Atoi(opts["port"])
	if err != nil {
		log.Fatal("mssql port number is improperly formatted")
	}
	debug, err := strconv.ParseBool(opts["debug"])
	if err != nil {
		log.Fatal("mssql debug value is improperly formatted")
	}

	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s", server, port, username, password)

	if debug {
		log.Printf(" connString:%s\n", connString)
	}

	// open the connection
	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open mssql connection failed:", err.Error())
	}
	defer conn.Close()

	// test the connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot connect to sql server:", err.Error())
	}

	return conn
}

func init() {
	client = connect()
}

func Client() *sql.DB {
	return client
}

func (ar *ArMsSql) SetKey(key string) {
	ar.ID = key
}

func (ar *ArMsSql) All(models interface{}, opts map[string]interface{}) (err error) {
	var limit int = 10 // per Orchestrate's documentation: 10 default, 100 max

	// set limit
	if opts["limit"] != nil {
		limit = opts["limit"].(int)
		if limit > 100 { // max limit is 100
			return errors.New("limit must be less than 100")
		}
	}

	//return mapResults(response.Results, models)
	return nil
}

func (ar *ArMsSql) Truncate() (numRowsDeleted int, err error) {
	return -1, nil
}

func (ar *ArMsSql) Find(id interface{}, out interface{}) error {
	//result, err := client.Get(ar.ModelName(), id.(string))

	//if result != nil {
	//err = result.Value(&out)
	//} else {
	//err = errors.New("record not found")
	//}

	return nil
}

func (ar *ArMsSql) DbSave() error {
	var err error

	//if ar.UpdatedAt != nil {
	//_, err = client.Put(ar.ModelName(), ar.ID, ar.Self())
	//} else {
	//_, err = client.PutIfAbsent(ar.ModelName(), ar.ID, ar.Self())
	//}

	return err
}

func (ar *ArMsSql) DbDelete() (err error) {
	//return client.Purge(ar.ModelName(), ar.ID)
	return nil
}

func (ar *ArMsSql) DbSearch(models interface{}) (err error) {
	var query, sort string
	//var response *c.SearchResults
	//query := r.Db(DbName()).Table(ar.Self().ModelName())

	// plucks
	//query = processPlucks(query, ar)

	// where conditions
	if query, err = processWhereConditions(ar); err != nil {
		return err
	}

	// aggregations
	//if query, err = processAggregations(query, ar); err != nil {
	//return err
	//}

	// order bys
	sort = processSorts(ar)

	// TODO: delete!
	log.Printf("DbSearch query: %s", query)

	// run search
	if sort == "" {
		//if response, err = client.Search(ar.ModelName(), query, 100, 0); err != nil {
		//return err
		//}
	} else {
		//if response, err = client.SearchSorted(ar.ModelName(), query, sort, 100, 0); err != nil {
		//return err
		//}
	}

	//return mapResults(response.Results, models)
	return nil
}

//func processPlucks(query r.Term, ar *ArRethinkDb) r.Term {
//if plucks := ar.Query().Plucks; plucks != nil {
//query = query.Pluck(plucks...)
//}

//return query
//}

func mapResults(orchestrateResults interface{}, models interface{}) (err error) {
	// now, map orchstrate's raw json to the desired active record type
	//modelsv := reflect.ValueOf(models)
	//if modelsv.Kind() != reflect.Ptr || modelsv.Elem().Kind() != reflect.Slice {
	//panic("models argument must be a slice address")
	//}
	//slicev := modelsv.Elem()
	//elemt := slicev.Type().Elem()

	//switch t := orchestrateResults.(type) {
	//case []c.KVResult:
	//for _, result := range t {
	//elemp := reflect.New(elemt)
	//if err = result.Value(elemp.Interface()); err != nil {
	//return err
	//}

	//slicev = reflect.Append(slicev, elemp.Elem())
	//}
	//case []c.SearchResult:
	//for _, result := range t {
	//elemp := reflect.New(elemt)
	//if err = result.Value(elemp.Interface()); err != nil {
	//return err
	//}

	//slicev = reflect.Append(slicev, elemp.Elem())
	//}
	//default:
	//return errors.New(fmt.Sprintf("Orchestrate Response Type Not Mapped: %v", t))
	//}

	//// assign mapped results to the caller's supplied array
	//modelsv.Elem().Set(slicev)

	//return err
	return nil
}

func processWhereConditions(ar *ArMsSql) (query string, err error) {
	var whereStmt, whereCondition string

	if len(ar.Query().WhereConditions) > 0 {
		for index, where := range ar.Query().WhereConditions {
			switch where.RelationalOperator {
			case EQ: // equal
				whereCondition = where.Key + ":" + fmt.Sprintf("%v", where.Value)
				//whereCondition = where.Key + ":" + where.Value.(string)
				//whereCondition = r.Row.Field(where.Key).Eq(where.Value)
			//case NE: // not equal
			//whereCondition = r.Row.Field(where.Key).Ne(where.Value)
			//case LT: // less than
			//whereCondition = r.Row.Field(where.Key).Lt(where.Value)
			//case LTE: // less than or equal
			//whereCondition = r.Row.Field(where.Key).Le(where.Value)
			//case GT: // greater than
			//// TODO: create function to set range based on type???
			//whereCondition = where.Key + ":[" + fmt.Sprintf("%v", where.Value) + " TO *]"
			//whereCondition = r.Row.Field(where.Key).Gt(where.Value)
			case GTE: // greater than or equal
				whereCondition = where.Key + ":[" + fmt.Sprintf("%v", where.Value) + " TO *]"
			//whereCondition = r.Row.Field(where.Key).Ge(where.Value)
			// case IN: // TODO: implement!!!!
			default:
				return query, errors.New(fmt.Sprintf("invalid comparison operator: %v", where.RelationalOperator))
			}

			if index == 0 {
				whereStmt = whereCondition
				//if where.LogicalOperator == NOT {
				//whereStmt = whereStmt.Not()
				//}
			} else {
				switch where.LogicalOperator {
				case AND:
					whereStmt = whereStmt + " AND " + whereCondition
					//whereStmt = whereStmt.And(whereCondition)
				case OR:
					whereStmt = whereStmt + " OR " + whereCondition
				//whereStmt = whereStmt.Or(whereCondition)
				////case NOT:
				////whereStmt = whereStmt.And(whereCondition).Not()
				default:
					whereStmt = whereStmt + " AND " + whereCondition
					//whereStmt = whereStmt.And(whereCondition)
				}
			}
		}

		// TODO: delete!!
		log.Printf("DbSearch whereStmt: %s", whereStmt)
		//query = query.Filter(whereStmt)
		//query = query.Filter(whereStmt)
	}

	return whereStmt, nil
}

//func processAggregations(query r.Term, ar *ArRethinkDb) (r.Term, error) {
//// sum
//if sum := ar.Query().Aggregations[SUM]; sum != nil {
//if len(sum) == 1 {
//query = query.Sum(sum...)
//} else {
//return query, errors.New(fmt.Sprintf("rethinkdb does not support summing more than one field at a time: %v", sum))
//}
//}

//// distinct
//if ar.Query().Distinct {
//query = query.Distinct()
//}

//return query, nil
//}

func processSorts(ar *ArMsSql) (sort string) {
	if len(ar.Query().OrderBys) > 0 {
		sort = ""

		for i, orderBy := range ar.Query().OrderBys {
			if i > 0 {
				sort += ","
			}

			sort += "value." + orderBy.Key + ":"

			switch orderBy.SortOrder {
			case DESC: // descending
				sort += "desc"
			default: // ascending
				sort += "asc"
			}
		}
	}

	return sort
}
