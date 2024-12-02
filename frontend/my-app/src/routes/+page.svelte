<script>
    import {number, z} from "zod"
   
    const emailSchema = z.string().email({message:"Invalid Email Address!"})
    
    let email = ''
    let number_of_days=''
    const validate = () => {
      try {
        emailSchema.parse(email)
        return true
      } catch (error) {
        if (error instanceof z.ZodError) {
          alert(error.errors[0].message)
          return false
        }
      }
    }
   
    const handleSubmit = async (event) => {
      event.preventDefault()
      if (validate()) {
        try {
          const response = await fetch('http://localhost:8090/saveData', {
            method: "POST",
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({ email, number_of_days })
          })
          if (response.ok) {
            alert("Data Saved!!")
          }
        } catch (error) {
          alert(error.message)
          console.log(error)
        }
      }
    }
   </script>
   
   <main>
    <form on:submit={handleSubmit}>
     <div class="form-container">
       <label for="email">Enter your email:</label>
       <input type="text" name="email" bind:value={email} />
       <label for="days">Number of Days:</label>
       <input type="text" bind:value={number_of_days}>
       <button type="submit">Submit!</button>
     </div>
    </form>
   </main>
   
   <style>
    main {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }
   
    .form-container {
      display: flex;
      flex-direction: column;
      gap: 10px;
      padding: 20px;
      border: 1px solid #ccc;
      border-radius: 8px;
      width: 300px;
      text-align: center;
    }
   
    input, button {
      width: 100%;
      padding: 8px;
    }
   </style>