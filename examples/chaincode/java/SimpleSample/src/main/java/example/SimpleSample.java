/*
Copyright DTCC 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

         http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package example;

import org.hyperledger.java.shim.ChaincodeBase;
import org.hyperledger.java.shim.ChaincodeStub;
import org.hyperledger.java.shim.ChaincodeResponse;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import com.google.protobuf.ByteString;

/**
 * <h1>Classic "transfer" sample chaincode</h1>
 * (java implementation of <A href="https://github.com/hyperledger/fabric/blob/master/examples/chaincode/go/chaincode_example02/chaincode_example02.go">chaincode_example02.go</A>)
 * @author Sergey Pomytkin spomytkin@gmail.com
 *
 */
public class SimpleSample extends ChaincodeBase {
	 private static Log log = LogFactory.getLog(SimpleSample.class);

	@Override
	public ByteString run(ChaincodeStub stub, String function, String[] args) {
		log.info("In run, function:"+function);
		
		switch (function) {
		case "init":
			init(stub, function, args);
			break;
		case "transfer":
			ByteString re = transfer(stub, args);	
			//System.out.println(re);
			return re;					
		case "put":
			for (int i = 0; i < args.length; i += 2)
				stub.putState(args[i], args[i + 1]);
			break;
		case "del":
			for (String arg : args)
				stub.delState(arg);
			break;
		case "query":
		    return query(stub, function, args);
			//break;
		default: 
			return transfer(stub, args);
		}
	 
		return ChaincodeResponse.Success("");
	}

	private ByteString transfer(ChaincodeStub stub, String[] args) {
		System.out.println("in transfer");
		if(args.length!=3){
			System.out.println("Incorrect number of arguments:"+args.length);
			String err = "{\"Error\":\"Incorrect number of arguments. Expecting 3: from, to, amount\"}";
			log.error(err);
			return ChaincodeResponse.Error("Incorrect number of arguments. Expecting 3: from, to, amount");
			//return "{\"Error\":\"Incorrect number of arguments. Expecting 3: from, to, amount\"}";
		}
		String fromName =args[0];
		String fromAm=stub.getState(fromName);
		String toName =args[1];
		String toAm=stub.getState(toName);
		String am =args[2];
		int valFrom=0;
		if (fromAm!=null&&!fromAm.isEmpty()){			
			try{
				valFrom = Integer.parseInt(fromAm);
			}catch(NumberFormatException e ){
				System.out.println("{\"Error\":\"Expecting integer value for asset holding of "+fromName+" \"}"+e);
				String err = "{\"Error\":\"Expecting integer value for asset holding of "+fromName+" \"}";
				log.error(err);
				return ChaincodeResponse.Error("Expecting integer value for asset holding of "+fromName);	
				//return "{\"Error\":\"Expecting integer value for asset holding of "+fromName+" \"}";		
			}		
		}else{
			String err = "{\"Error\":\"Failed to get state for " +fromName + "\"}";
			log.error(err);
			return ChaincodeResponse.Error("Failed to get state for " +fromName);	
			//return "{\"Error\":\"Failed to get state for " +fromName + "\"}";
		}

		int valTo=0;
		if (toAm!=null&&!toAm.isEmpty()){			
			try{
				valTo = Integer.parseInt(toAm);
			}catch(NumberFormatException e ){
				e.printStackTrace();
				String err = "{\"Error\":\"Expecting integer value for asset holding of "+toName+" \"}";
				log.error(err);
				return ChaincodeResponse.Error("Expecting integer value for asset holding of "+toName);
				//return "{\"Error\":\"Expecting integer value for asset holding of "+toName+" \"}";		
			}		
		}else{
			String err = "{\"Error\":\"Failed to get state for " +toName + "\"}";
			log.error(err);
			return ChaincodeResponse.Error("Failed to get state for " +toName);
			//return "{\"Error\":\"Failed to get state for " +toName + "\"}";
		}
		
		int valA =0;
		try{
			valA = Integer.parseInt(am);
		}catch(NumberFormatException e ){
			e.printStackTrace();
			String err = "{\"Error\":\"Expecting integer value for amount \"}";
			log.error(err);
			return ChaincodeResponse.Error("Expecting integer value for amount");
			//return "{\"Error\":\"Expecting integer value for amount \"}";
		}		
		if(valA>valFrom){
			String err = "{\"Error\":\"Insufficient asset holding value for requested transfer amount \"}";
			log.error(err);
			return ChaincodeResponse.Error("Insufficient asset holding value for requested transfer amount");
			//return "{\"Error\":\"Insufficient asset holding value for requested transfer amount \"}";
		}
			
		valFrom = valFrom-valA;
		valTo = valTo+valA;
		System.out.println("Transfer "+fromName+">"+toName+" am='"+am+"' new values='"+valFrom+"','"+ valTo+"'");
		stub.putState(fromName,""+ valFrom);
		stub.putState(toName, ""+valTo);		

		System.out.println("Transfer complete");

		return ChaincodeResponse.Success("");
		
	}

	public ByteString init(ChaincodeStub stub, String function, String[] args) {
		if(args.length!=4){
			String err = "{\"Error\":\"Incorrect number of arguments. Expecting 4\"}";
			log.error(err);
			return ChaincodeResponse.Error("Incorrect number of arguments. Expecting 4");
			//return "{\"Error\":\"Incorrect number of arguments. Expecting 4\"}";
		}
		try{
			int valA = Integer.parseInt(args[1]);
			int valB = Integer.parseInt(args[3]);
			stub.putState(args[0], args[1]);
			stub.putState(args[2], args[3]);		
		}catch(NumberFormatException e ){
			String err = "{\"Error\":\"Expecting integer value for asset holding\"}";
			log.error(err);
			return ChaincodeResponse.Error("Expecting integer value for asset holding");
			//return "{\"Error\":\"Expecting integer value for asset holding\"}";
		}		
		return ChaincodeResponse.Success("");
	}

	
	@Override
	public ByteString query(ChaincodeStub stub, String function, String[] args) {
		if(args.length!=1){
			String err = "{\"Error\":\"Incorrect number of arguments. Expecting name of the person to query\"}";
			log.error(err);
			return ChaincodeResponse.Error("Incorrect number of arguments. Expecting name of the person to query");
			//return "{\"Error\":\"Incorrect number of arguments. Expecting name of the person to query\"}";
		}
		String am =stub.getState(args[0]);
		if (am!=null&&!am.isEmpty()){
			try{
				int valA = Integer.parseInt(am);
				String res = "{\"Name\":\"" + args[0] + "\",\"Amount\":\"" + am + "\"}";
				log.info(res);
				return ChaincodeResponse.Success(am);
				//return  "{\"Name\":\"" + args[0] + "\",\"Amount\":\"" + am + "\"}";
			}catch(NumberFormatException e ){
				String err = "{\"Error\":\"Expecting integer value for asset holding\"}";
				log.error(err);
				return ChaincodeResponse.Error("Expecting integer value for asset holding");
				//return "{\"Error\":\"Expecting integer value for asset holding\"}";		
			}		}else{
				String err = "{\"Error\":\"Failed to get state for " + args[0] + "\"}";
				log.error(err);
				return ChaincodeResponse.Error("Failed to get state for " + args[0]);
				//return "{\"Error\":\"Failed to get state for " + args[0] + "\"}";
		}
		

	}

	@Override
	public String getChaincodeID() {
		return "simplesample:0";
	}

	public static void main(String[] args) throws Exception {
		new SimpleSample().start(args);
	}


}
