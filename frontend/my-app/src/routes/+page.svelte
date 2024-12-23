<script>
	import {z} from "zod"
	let email = ''
	let numberOfDays = ''
	let error = ''
	
	const emailSchema = z.string().email({message: "Kindly enter a valid email!!"})
	const daysSchema = z.number().min(1, {message: "Days must be at least 1"})
	
	const validate = () => {
	  try {
		emailSchema.parse(email)
		daysSchema.parse(Number(numberOfDays))
		return true
	  } catch (err) {
		if (err instanceof z.ZodError) {
		  error = err.errors[0].message
		}
		return false
	  }
	}
	
	const handleSubmit = async () => {
  error = ''
  if (validate()) {
    try {
      const response = await fetch('http://localhost:8090/email', {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          email,
          days: Number(numberOfDays)  // Changed from numberOfDays to days
        })
      })
      
      if (response.ok) {
        alert("Response Stored!")
        email = ''
        numberOfDays = ''
      } else {
        const errorText = await response.text()
        error = `Server error: ${errorText}`
      }
    } catch (err) {
      console.error('Fetch error:', err)
      error = "Connection failed. Please try again."
    }
  }
}
   </script>
   
   <div class="container">
	<link rel="apple-touch-icon" href="/custom_icon.png"/>
	 <form on:submit|preventDefault={handleSubmit}>
	   <div class="form-group">
		 <label for="email">Email</label>
		 <input
		   type="email"
		   id="email"
		   bind:value={email}
		   placeholder="Enter your email"
		   required
		 />
	   </div>
	   <div class="form-group">
		 <label for="days">Number of Days</label>
		 <input
		   type="number"
		   id="days"
		   bind:value={numberOfDays}
		   placeholder="Enter number of days"
		   min="1"
		   required
		 />
	   </div>
	   {#if error}
		 <div class="error">{error}</div>
	   {/if}
	   <button type="submit">Submit</button>
	 </form>
   </div>
   
   <style>
   .container {
	 display: flex;
	 justify-content: center;
	 align-items: center;
	 height: 100vh;
   }
   
   form {
	 width: 300px;
	 padding: 20px;
	 border: 1px solid #ccc;
	 border-radius: 8px;
	 box-shadow: 0 2px 4px rgba(0,0,0,0.1);
   }
   
   .form-group {
	 margin-bottom: 15px;
   }
   
   label {
	 display: block;
	 margin-bottom: 5px;
   }
   
   input {
	 width: 100%;
	 padding: 8px;
	 border: 1px solid #ddd;
	 border-radius: 4px;
   }
   
   .error {
	 color: red;
	 margin-bottom: 10px;
   }
   
   button {
	 width: 100%;
	 padding: 10px;
	 background-color: #007bff;
	 color: white;
	 border: none;
	 border-radius: 4px;
	 cursor: pointer;
   }
   
   button:hover {
	 background-color: #0056b3;
   }
   </style>